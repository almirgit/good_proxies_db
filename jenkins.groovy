pipeline {
    agent any
    environment {
        GO_VERSION = 'golang:1.22.0'  // Docker image for Go
        GOCACHE = "${WORKSPACE}/go-build-cache"  // Set GOCACHE to a writable directory
    }
    stages {
        stage('Checkout') {
            steps {
                //dir('src') {  // Specify the existing working directory
                  git branch: 'main',
                      url: 'git@github.com:almirgit/good_proxies_db.git',
                      credentialsId: 'jenkins_id_rsa'
                //}
            }
        }        
        
        stage('Build') {
            steps {
                script {
                    docker.image("${GO_VERSION}").inside {
                        sh 'mkdir -p $GOCACHE' // Ensure the cache directory exists
                        sh 'go env' // Verify Go environment variables
                        sh 'rm -f go.mod && go mod init good_proxies_db' // Make go.mod file
                        sh 'go mod tidy' // Ensure dependencies are installed
                        //sh 'chmod -R 777 /.cache' // Fix permissions
                        //sh 'GOROOT=`pwd`/.. go build -o good_proxies_db .'  // Build the Go binary
                        sh 'go build -o good_proxies_db .'  // Build the Go binary
                    }
                }
            }
        }
        
        //stage('Test') {
        //    steps {
        //        script {
        //            docker.image("${GO_VERSION}").inside {
        //                sh 'go test ./...'  // Run Go tests
        //            }
        //        }
        //    }
        //}
        
        stage('Archive Artifact') {
            steps {
                archiveArtifacts artifacts: 'good_proxies_db', allowEmptyArchive: true
            }
        }

        stage('Deploy binary') {
            steps {
                sh('scp -o StrictHostKeyChecking=no good_proxies_db almir@fra1.koderacloud.net:/tmp')
                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo cp /tmp/good_proxies_db /usr/local/bin"')
            }
        }

        stage('Systemd setup and start') {
            steps {
                sh('scp -o StrictHostKeyChecking=no systemd/good_proxies_db.service almir@fra1.koderacloud.net:/tmp')
                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo mkdir -p /var/log/freeproxy && sudo chown freeproxy:freeproxy /var/log/freeproxy"')
                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo cp /tmp/good_proxies_db.service /usr/lib/systemd/system/good_proxies_db.service"')
                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo systemctl daemon-reload"')
                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo systemctl restart good_proxies_db.service"')
            }
        }
    }
}
