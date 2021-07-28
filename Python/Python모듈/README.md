# 모듈

* 목차
  * 모듈과 패키지
  * 가상환경

# 모듈과 패키지

## 모듈과 패키지

* 모듈
  * 특정 기능을 가진 .py단위로 작성한 것
* 패키지
  * 특정 기능과 관련된 여러 모듈의 집합
  * 패키지 안에는 또 다른 서브 패키지를 포함
* 라이브러리 > 패키지 > 모듈

## 파이썬 표준 라이브러리 (Python Standard Library)

* 파이썬에 기본적으로 설치된 모듈과 내장함수
  * https://docs.python.org/3/library
* random.py, math.py 등

## 외부 라이브러리 (3rd 파티 라이브러리)

* pip install 로 설치하는 라이브러리들

## 파이썬 패키지 관리자(pip)

* PyPI(Python Package Index)에 저장된 외부 패키지들을 설치하도록 도와주는 패키지 관리 시스템
  * pypi.org

### 명령어

* 패키지 설치

  ```shell
  $ pip install pkg==1.0.1
  ```

  * 최신버전 / 특정버전 / 최소버전을 명시하여 설치할 수 있음 (디폴트가 최신)
  * 이미 설치되어 있는 경우 이미 설치되어 있음을 알리고 아무것도 하지 않음

* 패키지 삭제

  ```bash
  $ pip uninstall pkg
  ```

  * pip 는 패키지를 업그레이드 하는 경우 과거 버전을 자동으로 지워줌

* 패키지 목록 및 특정 패키지 정보

  ```shell
  # 목록
  $ pip list
  # 정보
  $ pip show
  ```

* 패키지 freeze

  * 설치된 패키지의 비슷한 목록을 만들지만, pip install에서 활용되는 형식으로 출력
  * 관습적으로 requirements.txt 라는 파일 형태로 만들어서 관리함

  ```shell
  $ pip freeze > requirements.txt
  
  ...
  clyent==1.2.2
  colorama @ file:///tmp/build/80754af9/colorama_1607707115595/work
  conda==4.10.1
  conda-build==3.21.4
  conda-content-trust @ file:///tmp/build/80754af9/conda-content-trust_1617045594566/work
  conda-package-handling @ file:///opt/concourse/worker/volumes/live/73497069-9b43-4ad9-50ec-1abb340e14eb/volume/conda-package-handling_1618262140058/work
  conda-repo-cli @ file:///tmp/build/80754af9/conda-repo-cli_1620168426516/work
  conda-token @ file:///tmp/build/80754af9/conda-token_1620076980546/work
  conda-verify==3.4.2
  ...
  ```

  ```shell
  # 설치
  $ pip install -r requirements.txt
  ```

* 다양한 파이썬 프로젝트에서 requirements.txt가 사용됨

# 가상환경

## 가상환경

* 파이썬 표준 라이브러리가 아닌 외부 패키지와 모듈을 사용하는 경우 모두 pip를 통해 설치해야함
* 복수의 프로젝트를 하는 경우 버전이 상이할 수 있음
  * 과거 외주 프로젝트 - django 버전 2.x
  * 신규 회사 프로젝트 - django 버전 3.x
* 이러한 경우 가상환경을 만들어 프로젝트별로 독립적인 패키지를 관리할 수 있음

## venv

* 가상 환경을 만들고 관리하는 데 사용하는 모듈(Python 3.5 부터)
* Python의 코어만 들고 오고 프로젝트에 필요한 라이브러리만 설치하는 **독립환경**으로 보면 됨
* 특정 디렉토리에 가상환경을 만들고, 고유한 파이썬 패키지 집합을 가질 수 있음.
  * 특정 폴더에 가상환경(패키지 집합 폴더 등)이 있고
  * 실행환경(예 - shell)에서 가상환경을 활성화 시켜
  * 해당 폴더에 있는 패키지를 관리/사용함

## 가상환경 생성

* 가상환경을 생성하면, 해당 디렉토리에 별도의 파이썬 패키지가 설치됨

  ```shell
  python -m venv <폴더명>
  # 일반적으로 폴더명은 venv로 관습적으로 지음
  ```


## 가상환경 활성화 / 비활성화

*  아래의 명령어를 통해 가상환경을 활성화

* <venv>는 가상환경을 포함하는 디렉토리의 경로

  ```shell
  # mac os
  $ source <venv>/bin/activate
  # win - bash
  $ source <venv>/Scripts/activate.bat
  
  # (venv) 표시가 생김
  
  # 비활성화
  $ deactivate
  ```

  * 활성화된 가상환경은 폴더에 국한되지 않고 작동!
  * 절대경로로 가상환경을 활성화할 수도 있음
  * 프로젝트 진입 후 바로 활성화할 수 있도록 각각의 프로젝트의 최상단에 위치

* 여러 개의 프로젝트를 할 경우 활성화된 가상환경을 비활성화해줘야 한다

## 가상환경 예시

* 동일 컴퓨터에서 별도의 가상환경
  * pj_A -> venv ->pip freeze -> pj_B -> venv pip install -r requirements.txt

# 모듈/패키지 활용하기

## 모듈 만들기 - check

* check.py에 짝수를 판별하는 함수(even)와 홀수를 판별하는 함수(odd)를 만들고 check 모듈을 활용

* 모듈을 활용하기 위해서는 import 문을 통해 가져옴

  ```python
  import check
  
  check.odd(3)  # True
  check.even(3) # False
  ```

  ```python
  from check import odd
  
  odd(5) # True
  even(4) # NameError
  ```

  ```python
  from check import *
  
  odd(5)  # True
  even(4) # True
  ```

## 패키지

* 패키지는 여러 모듈/하위 패키지로 구조화
  * 활용 예시 : package.module
* 모든 폴더에는 \_\_init\_\_.py를 만들어 패키지로 인식
  * 파이썬 3.3부터는 해당 파일이 없어도 되지만, 하위버전 호환 및 프레임워크 등에서도의 동작을 위해 생성하는 것을 권장

## 패키지 만들기

* 수학과 통계 기능이 들어간 패키지를 아래와 같이 구성

  * math의 tools : 자연상수 e, 원주율 pi, my_max
  * statistics의 tools: mean함수 등

* 아래와 같이 활용

  ```python
  # 이름이 겹침
  from mypackage.math import tools 
  from mypackage.statistics import tools as statistics_tools
  
  tools.my_max([1, 2])
  statistics_tools.mean([1, 2])