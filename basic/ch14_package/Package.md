# 패키지

코드를 묶는 가장 큰 단위

### 기능이 필요할 때 패키지를 먼저 찾아볼것

직접 코딩말고 패키지를 뒤져라!!

1. [표준 패키지](https://golang.org/pkg)에서 찾는 기능이 있는지 확인
2. [인기 패키지](https://github.com/avelino/awesome-go)에 있는지 확인


### GO 빌드 시 사용하는 패키지는 다음의 경로에서 찾는다

    1. Go 설치 경로의 기본 패키지
    2. GOPATH/pkg 폴더에서 설치된 패키지
    3. 현재 모듈 아래 있는 패키지 확인

### 패키지 제작 시 권고 사항

1. 패키지명은 가독성 있게
2. 모든 문자는 소문자


### 패키지를 import 할 때

**패키지 초기화**

    - 패키지 내 사용 전역변수 초기화
    - 패키지에서 init() 함수 실행