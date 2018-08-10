pipeline {
  agent any
  stages {
    stage('checkout') {
      steps {
        git(url: 'https://github.com/xjcloudy/snowflake.git', branch: 'master')
      }
    }
  }
}