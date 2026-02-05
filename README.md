# 📁 Go Standard Project Layout 정리

이 문서는 Go 커뮤니티에서 널리 사용되는 project-layout 패턴을 기반으로
Go 프로젝트 구조와 각 디렉토리의 역할을 한눈에 볼 수 있도록 정리한 문서이다.

====================================================

🎯 목적

- 코드 역할을 명확히 분리
- 확장 가능한 구조 유지
- 실무 및 협업에 적합한 구조 설계

====================================================

📂 기본 구조 개요

```plain text
project/
├── cmd/
├── internal/
├── pkg/
├── configs/
├── build/
├── deployments/
├── scripts/
├── test/
└── go.mod
```

핵심 디렉토리:
- cmd
- internal
- pkg

====================================================

✅ cmd/ — 실행 진입점 (Entry Point)

역할:
- 프로그램의 시작점(main.go)을 모아두는 폴더
- 실행 파일 하나당 하나의 폴더를 가지는 것이 원칙

예시 구조:

```plain text
cmd/
├── api/
│   └── main.go
├── worker/
│   └── main.go
└── cli/
    └── main.go
```

특징:
- main.go는 비즈니스 로직을 거의 가지지 않음
- 내부 코드를 조립(wiring)하는 역할만 수행

====================================================

✅ internal/ — 프로젝트 전용 핵심 코드

역할:
- 해당 프로젝트 내부에서만 사용하는 코드
- Go 컴파일러가 외부 import를 자동으로 차단

예시 구조:

```plain text
internal/
├── user/
├── order/
├── payment/
└── service/
```

또는 소규모 프로젝트 예시:

```plain text
internal/lotto/
  ├── controller/
  ├── domain/
  ├── service/
  └── view/
```
장점:
- 의존성 관리 명확
- 대형 프로젝트에 필수
- 실수 방지

====================================================

✅ pkg/ — 재사용 가능한 코드

역할:
- 외부 프로젝트에서도 import 가능한 라이브러리 코드

예시:

```plain text
pkg/
├── logger/
├── config/
└── database/
```

사용 기준:

프로젝트 전용 코드 → internal/
여러 프로젝트 공용 → pkg/

====================================================

📁 기타 디렉토리

configs/
- 설정 파일 관리 (yaml, json 등)

build/
- 빌드 관련 스크립트

deployments/
- Docker, Kubernetes 설정

scripts/
- 자동화 스크립트

test/
- 통합 테스트

====================================================

🧠 설계 철학

1) main은 가볍게 유지
   - 실행만 담당
   - 로직은 포함하지 않음

2) 패키지는 작고 명확하게 구성
   - 깊은 계층 구조 지양
   - 기능 중심 분리

   예시:
   internal/user
   internal/order

3) 의존성 흐름 단순화

   cmd → internal → pkg

   역방향 의존성은 지양

====================================================

📊 Java(Spring) 구조 비교

```plain text
Spring 구조            Go Project Layout
-----------------------------------------
Application         → cmd/app/main.go
Controller          → internal/...
Service             → internal/...
Repository          → internal/...
Util                → pkg/...


====================================================

🚀 소규모 프로젝트 예시

lotto-go/
├── go.mod
├── cmd/
│   └── lotto/
│       └── main.go
└── internal/
    └── lotto/
        ├── controller/
        ├── domain/
        ├── service/
        └── view/

====================================================

✅ 핵심 요약

- cmd : 실행 진입점
- internal : 프로젝트 핵심 코드
- pkg : 재사용 라이브러리
- main.go는 항상 얇게 유지
- 기능 중심 패키지 분리

====================================================
```

📌 결론

project-layout 구조는:

✔ 실무 친화적
✔ 유지보수 쉬움
✔ 확장성 뛰어남

Go 프로젝트의 사실상 표준 구조로 활용된다.