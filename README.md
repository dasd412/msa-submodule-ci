MSA에서 protobuf 서브 모듈을 사용하는 상황이라고 가정하자.

만약, 한 서비스 A에서 서브 모듈을 업데이트한다고 해보자.

그러면 그 서브 모듈을 의존하고 있는, 또 다른 서비스 B는
 서비스 A의 업데이트로 인해 악영향을 받을 수 있다. 

특히 protobuf 같은 서비스 사이의 규약이 업데이트될 때 문제가 된다.

시나리오를 풀어쓰면,

1. 서비스 A와 서비스 B는 서브 모듈 내 c.proto를 둘다 사용하고 있다.
2. 서비스 A에서 서브 모듈 내 c.proto를 갱신한다.
3. 서비스 B는 그대로인채, 서비스 A만 배포된다면 proto 메시지를 보는 게 다르기 때문에 에러가 발생한다.

이러한 문제를 방지하기 위해, 서브 모듈의 버전이 최신인지 체크하기 위한 CI를 테스트해보려 한다.

***

# 절차
1. git 메인 프로젝트 추가 및 구성
2. git submodule 리포지토리 추가 및 구성
3. git 메인 프로젝트에 .gitmodules 구성
4. git 메인 프로젝트에서 다음 명령어 실행
```
git clone https://github.com/dasd412/msa-submodule-protobuf.git protobuf
git submodule add https://github.com/dasd412/msa-submodule-protobuf.git protobuf
git add .gitmodules protobuf
git commit -m "Add protobuf as submodule"
git push origin main
```
5. 서브 모듈 내 있는 .proto 파일을 이용해 pb.go 파일을 생성
6. 

***
