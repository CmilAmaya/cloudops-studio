pipeline {
    agent any

    environment {
        PATH = "/opt/homebrew/bin:/usr/local/bin:/usr/bin:/bin"
        BACKEND_IMAGE = "todo-backend"
        FRONTEND_IMAGE = "todo-frontend"
    }

    stages {
        stage('Checkout') {
            steps {
                echo "Clonando repositorio..."
                git url: 'https://github.com/CmilAmaya/cloudops-studio.git', branch: 'main'
            }
        }

        stage('Hadolint - Backend Dockerfile') {
            steps {
                sh '''
                docker run --rm -i hadolint/hadolint < backend/Dockerfile
                '''
            }
        }

        stage('Hadolint - Frontend Dockerfile') {
            steps {
                sh '''
                docker run --rm -i hadolint/hadolint < frontend/Dockerfile
                '''
            }
        }

        stage('Build Backend Docker Image') {
            steps {
                echo "Construyendo imagen del backend..."
                sh "docker build -t ${env.BACKEND_IMAGE} ./backend"
            }
        }

        stage('Build Frontend Docker Image') {
            steps {
                echo "Construyendo imagen del frontend..."
                sh "docker build -t ${env.FRONTEND_IMAGE} ./frontend"
            }
        }

        stage('Trivy Scan - Backend Image') {
            steps {
                sh '''
                docker run --rm \
                -v /var/run/docker.sock:/var/run/docker.sock \
                aquasec/trivy image ${BACKEND_IMAGE}
                '''
            }
        }

        stage('Trivy Scan - Frontend Image') {
            steps {
                sh '''
                docker run --rm \
                -v /var/run/docker.sock:/var/run/docker.sock \
                aquasec/trivy image ${FRONTEND_IMAGE}
                '''
            }
        }

        stage('Dive - Backend Image') {
            steps {
                sh '''
                docker run --rm \
                -v /var/run/docker.sock:/var/run/docker.sock \
                wagoodman/dive:latest ${BACKEND_IMAGE} --ci
                '''
            }
        }

        stage('Start Services with Docker Compose') {
            steps {
                echo "Levantando servicios con Docker Compose..."
                sh 'docker-compose down -v || true'
                sh "docker-compose up -d"
            }
        }

        stage('Run Migrations') {
            steps {
                echo "Ejecutando migraciones..."
                sh "docker-compose run --rm migrations"
            }
        }

        stage('Run Backend Tests') {
            steps {
                dir('backend') {
                    sh 'go test ./...'
                }
            }
        }

        stage('Run Frontend Smoke Test') {
            steps {
                echo "Verificando que frontend responde..."
                sh 'curl -f http://localhost:3000 || exit 1'
            }
        }
    }

    post {
        always {
            echo "Limpiando todos los contenedores..."
            sh "docker-compose down -v"
        }
        success {
            echo "Pipeline completado correctamente!"
        }
        failure {
            echo "Pipeline falló."
        }
    }
}
