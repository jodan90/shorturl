# Go 프로젝트

이 프로젝트는 Go 언어로 만든 URL 단축기입니다. 긴 URL을 입력하면 짧은 URL로 변환하여 제공하며, 변환된 URL을 통해 원래 URL로 리다이렉션할 수 있습니다.

## 주요 기능
- **URL 단축**: 긴 URL을 입력하면 짧은 URL을 생성.
- **리다이렉션**: 생성된 짧은 URL을 통해 원래의 긴 URL로 리다이렉트.
- **모던한 UI**: Bootstrap을 사용하여 깔끔하고 반응형 디자인 적용.
- **URL 저장**: 실행 중인 동안 메모리에 URL을 저장.


## 실행 방법

1. **레포지토리 클론**:
    ```bash
    git clone https://github.com/jodan90/shorturl.git
    cd url-shortener-go
    ```

2. **Go 모듈 초기화**:
    ```bash
    go mod init url-shortener
    ```

3. **애플리케이션 실행**:
    ```bash
    go run main.go
    ```

4. **브라우저에서 애플리케이션 접근**:
    - `http://localhost:8080`로 접속하여 URL 단축기를 사용합니다.
    - 긴 URL을 입력하면 단축된 URL을 얻을 수 있습니다.

## 의존성
- [Bootstrap](https://getbootstrap.com): 모던한 UI 컴포넌트를 위해 사용.
- Go 1.16+ 버전이 필요.

## 추가 개발 계획
- URL을 데이터베이스에 저장하여 영구적인 단축 URL 제공.
- 단축 URL 클릭 횟수 통계 제공.
- URL 만료 기능 및 사용자 정의 단축 URL 지원.

## 라이센스
이 프로젝트는 [MIT License](LICENSE) 하에 오픈소스 프로젝트로 제공됩니다.

