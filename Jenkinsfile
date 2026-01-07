pipeline {
    agent any

    environment {
        PATH+EXTRA = "/usr/local/bin"
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

        stage('Start Services with Docker Compose') {
            steps {
                echo "Levantando servicios con Docker Compose..."
                sh "docker-compose up -d --build"
            }
        }

        stage('Run Migrations') {
            steps {
                echo "Ejecutando migraciones..."
                sh "docker-compose run --rm migrations"
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
