pipeline {
    agent any

    environment {
        DOCKER_IMAGE = 'todo-cli'
    }

    stages {
        stage('Build') {
            steps {
                script {
                    // Clean the workspace
                    sh 'go clean -modcache'
                    
                    // Build the Go application
                    sh 'go build -o myapp .'
                }
            }
        }

        stage('Test') {
            steps {
                script {
                    // Run tests
                    sh 'go test ./...'
                }
            }
        }

        stage('Docker') {
            steps {
                script {
                    // Build the Docker image
                    sh "docker build -t ${DOCKER_IMAGE}:${BUILD_ID} ."
                }
            }
        }

        stage('Post Build') {
            steps {
                script {
                    // Optionally push the Docker image to a registry
                    // Uncomment the following line if you want to push the image
                    // sh "docker push ${DOCKER_IMAGE}:${DOCKER_TAG}"

                    // Clean up the workspace
                    sh 'go clean'
                }
            }
        }
    }

    post {
        always {
            // Archive the build artifacts (optional)
            archiveArtifacts artifacts: 'myapp', fingerprint: true

            // Clean up Docker images (optional)
            sh "docker rmi ${DOCKER_IMAGE}:${DOCKER_TAG} || true"
        }
        success {
            echo 'Build and Docker image creation succeeded!'
        }
        failure {
            echo 'Build or tests failed.'
        }
    }
}
