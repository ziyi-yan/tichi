package chingwei

import (
	"github.com/sirupsen/logrus"
	"k8s.io/test-infra/prow/github"
)

type githubClient interface {
	FindIssues(query, sort string, asc bool) ([]github.Issue, error)
}

func Reproducing(log *logrus.Entry, ghc githubClient) error {
	log.Info("Staring search pull request.")
	query := `repo:"tidb" label:"status/needs-reproduction"`

	issues, err := ghc.FindIssues(query, "", false)
	if err != nil {
		return err
	}

	for _, issue := range issues {
		log.WithField("issue", issue).Info("Get issue success.")
		
		// TODO parse minimal reproduce step and version from issue
		sql := "minimal reproduce step"
		mysqlVersion := "8.0"
		tidbVersion := "v5.0.0-rc"
		
		// TODO try send a SQL query to tidb with a specific version
		tidbDSN, tidbCleanup, err := PrepareTiDB(tidbVersion)
		mysqlDSN, mysqlCleanup, err := PrepareMySQL(mysqlVersion)
	
		// TODO: do reproduce by connecting to `dsn`
		tidbOutput, tidbErr := Reproduce(tidbDSN, sql)
		mysqlOutput, mysqlErr := Reproduce(mysqlDSN, sql)

		// do cleanup
		tidbCleanup()
		mysqlCleanup()
	}

	

	return nil
}

// PrepareTiDB start a tidb server with specific version.
// It returns a function to destory and cleanup the related resources.
func PrepareTiDB(version string) (string, func () {}, error) {

}