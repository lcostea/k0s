package cmd

import (
	"bytes"
	"encoding/base64"

	"github.com/cloudflare/cfssl/log"
	"github.com/k0sproject/k0s/internal/util"
	"github.com/k0sproject/k0s/pkg/certificate"
	"github.com/k0sproject/k0s/pkg/constant"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"html/template"
	"io/ioutil"
	"os"
	"path"
)

func init() {
	kubeconfigCreateCmd.Flags().StringVar(&groups, "groups", "", "Specify groups")
	kubeconfigCmd.AddCommand(kubeconfigCreateCmd)
	kubeconfigCmd.AddCommand(kubeConfigAdminCmd)
}

var (
	groups string

	userKubeconfigTemplate = template.Must(template.New("kubeconfig").Parse(`
apiVersion: v1
clusters:
- cluster:
    server: {{.JoinURL}}
    certificate-authority-data: {{.CACert}}
  name: k0s
contexts:
- context:
    cluster: k0s
    user: {{.User}}
  name: k0s
current-context: k0s
kind: Config
preferences: {}
users:
- name: {{.User}}
  user:
    client-certificate-data: {{.ClientCert}}
    client-key-data: {{.ClientKey}}
`))

	// kubeconfigCmd creates new certs and kubeConfig for a user
	kubeconfigCmd = &cobra.Command{
		Use:   "kubeconfig [command]",
		Short: "Manage user access",
		RunE: func(cmd *cobra.Command, args []string) error {
			// user command does nothing
			return kubeconfigCreateCmd.Usage()
		},
	}

	kubeconfigCreateCmd = &cobra.Command{
		Use:   "create [username]",
		Short: "Create a kubeconfig for a user",
		Long: `Create a kubeconfig with a signed certificate and public key for a given user (and optionally user groups)
Note: A certificate once signed cannot be revoked for a particular user`,
		Example: `	Command to create a kubeconfig for a user:
	CLI argument:
	$ k0s kubeconfig create [username]

	optionally add groups:
	$ k0s kubeconfig create [username] --groups [groups]`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// Disable logrus and cfssl logging for user commands to avoid printing debug info to stdout
			logrus.SetLevel(logrus.FatalLevel)
			log.Level = log.LevelFatal

			if len(args) == 0 {
				return errors.New("Username is mandatory")
			}
			var username = args[0]
			clusterConfig, err := ConfigFromYaml(cfgFile)
			if err != nil {
				return err
			}
			var config = constant.GetConfig(dataDir)

			caCert, err := ioutil.ReadFile(path.Join(config.CertRootDir, "ca.crt"))
			if err != nil {
				return errors.Wrapf(err, "failed to read cluster ca certificate, is the control plane initialized on this node?")
			}

			caCertPath, caCertKey := path.Join(config.CertRootDir, "ca.crt"), path.Join(config.CertRootDir, "ca.key")

			if err != nil {
				return err
			}

			userReq := certificate.Request{
				Name:   username,
				CN:     username,
				O:      groups,
				CACert: caCertPath,
				CAKey:  caCertKey,
			}
			certManager := certificate.Manager{
				K0sVars: config,
			}
			userCert, err := certManager.EnsureCertificate(userReq, "root")
			if err != nil {
				return err
			}

			data := struct {
				CACert     string
				ClientCert string
				ClientKey  string
				User       string
				JoinURL    string
			}{
				CACert:     base64.StdEncoding.EncodeToString(caCert),
				ClientCert: base64.StdEncoding.EncodeToString([]byte(userCert.Cert)),
				ClientKey:  base64.StdEncoding.EncodeToString([]byte(userCert.Key)),
				User:       username,
				JoinURL:    clusterConfig.Spec.API.APIAddress(),
			}

			var buf bytes.Buffer

			err = userKubeconfigTemplate.Execute(&buf, &data)
			if err != nil {
				return err
			}
			os.Stdout.Write(buf.Bytes())
			return nil
		},
	}

	kubeConfigAdminCmd = &cobra.Command{
		Use:   "admin [command]",
		Short: "Manage user access",
		Long:  "Command dumps admin kubeconfig.",
		Example: `	$ k0s kubeconfig admin > kubeconfig
	$ export KUBECONFIG=kubeconfig
	$ kubectl get nodes`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if util.FileExists(k0sVars.AdminKubeconfigConfigPath) {
				content, err := ioutil.ReadFile(k0sVars.AdminKubeconfigConfigPath)
				if err != nil {
					log.Fatal(err)
				}
				os.Stdout.Write(content)
			} else {
				return errors.Errorf("failed to read admin config, is the control plane initialized on this node?")
			}
			return nil
		},
	}
)
