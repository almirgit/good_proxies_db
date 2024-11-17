pipeline {
    agent any
    triggers {
        pollSCM('H/2 * * * *') // Poll SCM every 2 minutes for changes
    }
    stages {
        stage('Demo Step') {
            steps {
                echo 'This is a dummy step demonstrating the pipeline execution.'
            }
        }
    }
}

//OK//pipeline {
//OK//    agent any
//OK//    environment {
//OK//        GO_VERSION = 'golang:1.22.0'  // Docker image for Go
//OK//        GOCACHE = "${WORKSPACE}/go-build-cache"  // Set GOCACHE to a writable directory
//OK//    }
//OK//    stages {
//OK//        stage('Checkout') {
//OK//            steps {
//OK//                //dir('src') {  // Specify the existing working directory
//OK//                  git branch: 'main',
//OK//                      url: 'git@github.com:almirgit/good_proxies_db.git',
//OK//                      credentialsId: 'jenkins_id_rsa'
//OK//                //}
//OK//            }
//OK//        }        
//OK//        
//OK//        stage('Build') {
//OK//            steps {
//OK//                script {
//OK//                    docker.image("${GO_VERSION}").inside {
//OK//                        sh 'mkdir -p $GOCACHE' // Ensure the cache directory exists
//OK//                        sh 'go env' // Verify Go environment variables
//OK//                        sh 'rm -f go.mod && go mod init good_proxies_db' // Make go.mod file
//OK//                        sh 'go mod tidy' // Ensure dependencies are installed
//OK//                        //sh 'chmod -R 777 /.cache' // Fix permissions
//OK//                        //sh 'GOROOT=`pwd`/.. go build -o good_proxies_db .'  // Build the Go binary
//OK//                        sh 'go build -ldflags "-X main.version=$(git describe --tags) -X main.commitSHA=$(git rev-parse --short HEAD)" -o good_proxies_db'
//OK//                    }
//OK//                }
//OK//            }
//OK//        }
//OK//        
//OK//        //stage('Test') {
//OK//        //    steps {
//OK//        //        script {
//OK//        //            docker.image("${GO_VERSION}").inside {
//OK//        //                sh 'go test ./...'  // Run Go tests
//OK//        //            }
//OK//        //        }
//OK//        //    }
//OK//        //}
//OK//        
//OK//        stage('Archive Artifact') {
//OK//            steps {
//OK//                archiveArtifacts artifacts: 'good_proxies_db', allowEmptyArchive: true
//OK//            }
//OK//        }
//OK//
//OK//        stage('Deploy binary') {
//OK//            steps {
//OK//                sh('scp -o StrictHostKeyChecking=no good_proxies_db almir@fra1.koderacloud.net:/tmp')
//OK//                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo cp -f /tmp/good_proxies_db /usr/local/bin"')
//OK//            }
//OK//        }
//OK//
//OK//        stage('Systemd setup and start') {
//OK//            steps {
//OK//                sh('scp -o StrictHostKeyChecking=no systemd/good_proxies_db.service almir@fra1.koderacloud.net:/tmp')
//OK//                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo mkdir -p /var/log/freeproxy && sudo chown freeproxy:freeproxy /var/log/freeproxy"')
//OK//                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo cp /tmp/good_proxies_db.service /usr/lib/systemd/system/good_proxies_db.service"')
//OK//                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo systemctl daemon-reload"')
//OK//                sh('ssh -o StrictHostKeyChecking=no almir@fra1.koderacloud.net "sudo systemctl restart good_proxies_db.service"')
//OK//            }
//OK//        }
//OK//    }
//OK//}
