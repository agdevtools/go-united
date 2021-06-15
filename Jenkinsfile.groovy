// Run on an agent where we want to use Go

node {
    def gitRepository = 'https://github.com/agdevtools/go-postgres.git/'
    def gitBranch = '*/master'
    def githubcreds = [
            $class      : 'UsernamePasswordMultiBinding',
            credentialsId : 'githubcreds',
            usernameVariable : 'GIT_USER',
            passwordVariable : 'GIT_PASS'
    ]

    stage('Clean Workspace') {
        cleanWs()
    }

    stage('Git Checkout') {
        withCredentials([githubcreds]){
            checkout([
                    $class      : 'GitSCM',
                    branches    : [[name:"${gitBranch}"]],
                    doGenerateSubModuleConfigurations : false,
                    extensions: [],
                    submoduleCfg: [],
                    userRemoteConfigs: [[credentialsId  : 'githubcreds',
                                         url            :"${gitRepository}"]]

            ])

        }

        // Ensure the desired Go version is installed
        def root = tool type: 'go', name: 'Go 1.16.4'

        // Export environment variables pointing to the directory where Go was installed
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
        }
    }

    stage('Build') {
        def root = tool type: 'go', name: 'Go 1.16.4'
        withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
            sh 'go version'
            sh 'go get -u golang.org/x/lint/golint'
            sh 'go build -o bin/main main/main.go'

        }
    }

    stage('Test') {

        sh 'echo running tests ...'

    }

    stage('Deploy to Heroku') {

        sh 'git checkout master'
        sh 'git fetch'
        sh 'git pull'
        sh 'git merge origin/add-feature'
        sh 'git commit --allow-empty -m "merge feature branch with master to deploy"'
        sh 'git push'

    }
}