pipeline {
    environment {
        QODANA_TOKEN = credentials('qodana-token')
    }
    agent {
        docker {
            args '''
                -v "${WORKSPACE}":/data/project
                --entrypoint=""
                '''
            image 'jetbrains/qodana-go'
        }
    }
    stages {
        stage('Qodana') {
            when {
                branch 'main'
                branch 'wip/create_product_structure'
            }
            steps {
                sh '''qodana'''
            }
        }
    }
}