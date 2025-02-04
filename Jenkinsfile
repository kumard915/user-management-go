pipeline {
    agent any

    environment {
        DOCKER_IMAGE = "user-management"
        DOCKER_REPO = "kumard915/user-management"
    }

    stages {
        stage('Checkout Code') {
            steps {
                git 'https://github.com/kumard915/user-management-go.git'
            }
        }

        stage('Install Dependencies') {
            steps {
                sh 'go mod tidy'
            }
        }

        stage('Build Go App') {
            steps {
                sh 'go build -o main'
            }
        }

        stage('Run Unit Tests') {
            steps {
                sh 'go test ./...'
            }
        }

        stage('Build Docker Image') {
            steps {
                sh 'docker build -t $DOCKER_IMAGE .'
            }
        }

        stage('Push to Docker Hub') {
            steps {
                withDockerRegistry([credentialsId: 'docker-hub-credentials', url: '']) {
                    sh 'docker tag $DOCKER_IMAGE $DOCKER_REPO:latest'
                    sh 'docker push $DOCKER_REPO:latest'
                }
            }
        }

        stage('Deploy Container') {
            steps {
                sh 'docker stop $DOCKER_IMAGE || true'
                sh 'docker rm $DOCKER_IMAGE || true'
                sh 'docker run -d -p 8080:8080 --name $DOCKER_IMAGE $DOCKER_REPO:latest'
            }
        }
    }

    post {
        success {
            echo "🚀 Deployment Successful!"
        }
        failure {
            echo "❌ Build Failed!"
        }
    }
}
