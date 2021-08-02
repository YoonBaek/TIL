# HTML

# 개발 환경 설정

* Visual Studio Code
  * Open in Browser
  * Name in Tag 등 익스텐션

# HTML & CSS

## Why Web?

* 웹을 왜 배워야 하는가
  * 웹 어플리케이션 개발을 하기 위해? 
  * SW 개발 방법 및 학습 과정을 익히기 위해서!
  * SW의 기반이되는 웹

## 현재 웹 표준

* 왜 표준을 지켜야 하는가

  * HTML은 브라우저에 나타내기 위한 도구인데 각 표준이 다를 경우 브라우저마다의 표준을 알아둬야 함

* W3C(World Wide Web)에서 표준을 관리

  * HTML5
  * 그러나 표준 인증을 거치는 개발과정이 너무 느림

* WHATWG

  * 웹 개발 조직 들의 집합
  * Apple, Google, Microsoft, Mozila
  * 모두 브라우저를 갖고 있는 업체
  * 자체적으로 표준을 조율함

* 현재는 W3C가 주도권을 잃고 WHATWG(왓위그)가 승리

* 누가누가 표준을 잘 지키나

  * html5test.com
    * Chrome 528
    * Opera 518
    * FireFox 492
    * ...
    * Internet Explorer 312

  * Can I use?
    * 기능을 사용할 수 있는지 브라우저 별로 보여주는 사이트

# HTML

hyper text markup language

## Hyper Text

* Hyper
  * 텍스트 등의 정보가 동일 선상에 있는 것이 아니라 다중으로 연결되어 있는 상태
* Hyper Text
  * 참조(hyper link)를 통해 사용자가 한 문서에서 다른 문서로 즉시 접근할 수 있는 텍스느
  * 가장 중요한 기술 2가지

## Markup

* 특정 텍스트에 역할을 부여하는 것
  * \<h1\> HTML \<\/h1\>
    * 너는 큰 제목이야
* 태그 등을 이용하여 문서나 데이터의 구조를 명시하는 언어
* 프로그래밍 언어와는 다르게 단순하게 데이터를 표현하기만 한다
* 대표적인 예 - HTML, Markdown

## HTML

* 웹 페이지를 작성하기 위한(구조를 잡기위한) 언어
* 웹 컨텐츠의 의미와 구조를 정의
* .html의 확장자를 사용

# HTML 기본 구조

## 기본 구조

```html
<!DOCTYPE html>
<html lang="ko">
  <head>
    <meta charset="UTF-8">
    <title>Document</title>
  </head>
  <body>

  </body>
</html>
```

* 전체를 html이 감싸고 있음
  * 문서의 root, 최상단을 뜻함
* **head 요소**
  * 문서 제목, 문자코드(인코딩)와 같이 해당 문서 정보를 담고 있으며, 브라우저에 나타나지 않는다.
  * CSS 선언 외부 로딩 파일 지정 등도 작성합니다.
  * 사용자에게 보이는 부분은 아니라는게 핵심
* **body 요소**
  * 브라우저 화면에 나타나는 정보로 실제 내용에 해당한다.
  * 대부분의 부분이 body에 해당

## Open Graph Protocol

메타 데이터를 표현하는 새로운 규약

* HTML 문서의 메타 데이터를 통해 문서의 정보를 전달
* 페이스북에서 만들었으며, 메타정보에 해당하는 제목, 설명 등을 쓸 수 있도록 정의
  * 카카오톡에서 url 공유했을 때 박스형태로 정보까지 나오는게 대표적인 예시
* metatags.io

## DOM 트리 구조

* Document Object Model

* 들여쓰기 별로 하나의 객체가 있는 것

  ```html
  <body>
      <h1>웹문서</h1>
      <ul>
          <li>HTML</li>
          <li>CSS</li>
      </ul>
  </body>
  ```

* 2칸, 4칸 공백을 설정할 수 있는데 최근에는 2칸 공백을 위주로 사용하는 추세
* DOM은 문서의 구조화된 표현(Structured Representation)을 제공하며, 프로그래밍 언어가 DOM 구조에 접근할 수 있는 방법을 제공하여 그들이 문서 구조, 스타일 내용 등을 변경할 수 있게 도움
* DOM은 동일한 문서를 표현하고, 저장하고, 조작하는 방법을 제공
* Web Page의 객체 지향 표현

## 요소 (element)

* HTML의 요소는 태그와 내용(contents)로 구성되어 있다.

  ```html
  <h1>contents</h1>
  <!--<여는태그>내용</닫는태그>-->
  ```

* 태그는 컨텐츠를 감싸는 것으로 그 정보의 성격과 의미를 정의
* 내용이 없는 태그들도 있음
  * br, hr, img, input, link, meta
* 요소는 중첩(nested)될 수 있음
  * 요소의 중첩을 통해 하나의 문서를 구조화
  * 여는 태그와 닫는 태그의 쌍을 잘 확인해야함
  * 오류를 반환하는 것이 아닌 그냥 레이아웃이 깨진 상태로 출력되기 때문에, 디버깅이 어려울 수 있음

## 속성 (attribute)

* 태그 별로 사용할 수 있는 속성은 다르다.

  ```html
  <a href="https://google.com"></a>
  <!--<태그 속성=속성값>
  띄어쓰기, 쌍따옴표 주의
  -->
  ```

* 속성을 통해 태그의 부가적인 정보를 설정할 수 있음

* 요소는 속성을 가질 수 있으며. 경로나 크기와 같은 추가적인 정보를 제공

* 요소의 시작 태그에 작성하며 보통 이름과 값이 하나의 쌍으로 존재

* 태그와 상관없이 사용 가능한 속성들(html global attribute)도 있음

## HTML Global Attribute

* 모든 HTML 요소가 공통으로 사용할 수 있는 속성 (몇몇 요소에는 아무 효과가 없을 수 있음)

  * id, class
  * hidden
  * lang
  * style
  * tabindex
  * title

* **MDN (mozila development networks)**에서 확인할 수 있음

  * 구글, ms w3c 등의 웹 표준도 모두 이 곳에 작성
  * 사실상의 공식 문서

* 코드로 확인

  ```html
  <!DOCTYPE html>
  <html>
  <head>
    <meta charset="UTF-8">
    <!--창의 제목 역할을 함-->
    <title>타이틀입니다.</title>
  <body>
    <h1>나의 첫번째 HTML</h1>
    <!--a 태그 : 링크를 연결할 수 있도록 함-->
    <a href="https://google.com">google로 고고!</a> 
  </body>
  </head>
  </html>
  ```

## 시멘틱 태그

* HTML5에서 의미론적 요소를 담은 태그의 등장. div를 의미론적으로 세분화한 것 div와 역할은 같다.

  ```html
  <div>
    <div></div>
  </div>
  <div>
    <div></div>
    <div></div>
  </div>
  ```

  의미론적으로 무슨 뜻인지 알기 어렵다

  * 이건 자리야.
    * 무슨 자리인데?
  * 자리.
    * 아오

* div를 세분화한 대표적인 태그들은 다음과 같다.

  * header: 문서 전체나 섹션의 헤더(머릿말 부분)
  * nav: 내비게이션
  * aside: 사이드에 위치한 공간, 메인 콘텐츠와 관련성이 적은 콘텐츠
  * section: 문서의 일반적인 구분, 컨텐츠의 그룹을 표현
  * article: 문서, 페이지, 사이트 안에서 독립적으로 구분되는 영역
  * footer: 문서 전체나 섹션의 푸터(마지막 부분)
  * 이건 헤더 자리고 저건 아티클 자리야.
    * ㅇㅋ 명확쓰

* 장점
  1. 읽기가 좋아진다.
  2. 검색엔진이 접근하기 좋아진다.

* 개발자 및 사용자 뿐만 아니라 검색엔진 등에 의미 있는 정보의 그룹을 태그로 표현
* 단순히 구역을 나누는 것 뿐만 아니라 '의미'를 가지는 태그들을 활용하기 위한 노력
* Non semantic 요소는 div, span 등이 있으며 h1, table 태그들도 시맨틱 태그로 볼 수 있음
* 요소의 의미가 명확해지기 때문에 코드의 가독성을 높이고 유지보수를 쉽게 함
* 검색엔진최적화(SEO)를 위해서 메타태크, 시맨틱 태그 등을 통한 마크업을 효과적으로 활용할 필요가 있다.

## 시맨틱 웹

* 웹 상에 존재하는 수많은 웹 페이지들에 메타데이터를 부여하여,
* 기존의 단순한 데이터의 집압이었던 웹페이지를 **'의미'**와 '관련성'을 가지는 거대한 데이터 베이스로 구축하고자 하는 발상

# HTML 문서 구조화

## 인라인 / 블록 요소

* 인라인 : 라인 전체 차지. 세로 공간을 갖고 있지 않아 전체를 활용
* 블록 : 자기 공간만큼만 차지. 세로 공간을 가짐.

## 그룹 컨텐츠

* \<p> : 문단
* \<hr> : 헤드라인
* \<ol>, \<ul> : 오더리스트, 언오더리스트
* \<pre>, \<blockquote> : 주석, 인용문
* \<div> : 텍스트를 그룹화

## 텍스트 관련 요소

* \<a> : 하이퍼링크 텍스트 생성
* \<b> vs \<strong> : 굵게하는 역할
  * b : 문서 표현상으로만 굵게 표시
  * strong : 굵게 표시함과 동시에 강조까지 함
    * 장애인들을 위한 웹 접근성에서 실제로 강하게 읽어주는 기능
* \<i> vs \<em>
* \<span>, \<br>, \<img> 
  * span : div와 역할이 같음
  * br : 줄바꿈
  * img : 이미지
* 기타 등등

## table

* 테이블 태그는 최근에는 직접적으로 사용하진 않음

## form

* \<form>은 서버로 데이터를 보내주는 역할을 함
  * 아이디 로그인, 검색, 등등...
* \<form>의 기본 속성
  * action : 어디로?
  * method : http method 선택
* 서버가 있어야지 동작할 수 있음

## input

* form 안에 들어감.
* 사용자로부터 입력을 받는 부분
* 다양한 타입을 가지고 있음
  * 버튼, 체크박스, 날짜, 이메일, 파일, 비밀번호 ...
* autofocus 기능 지원
  * 페이지에 접속하면 자동으로 커서를 인풋단에 놓아줌
* option : 드랍다운 기능
  * required 필수 선택 역할
  * disabled 해당 목록 비활성화
* submit : 입력 값을 서버로 전송하는 타입
  * value로 버튼 이름을 설정 가능
* label : 라벨을 키형태로 다른 단위에 연결해줌

# CSS

* Cascading Style Sheet
* html을 기반으로 스타일 / 레이아웃 등을 결정하며 클라이언트에게 어떻게 보일지를 다룸
* HTML 문서를 표시하는 방법을 지정하는 언어

## 들어가기 전에

* 마음대로 되지 않는게 css
* 만질 수록 망하기도 함 ㅋㅋ

## CSS 구문

* 선택자로 열게 됨

* 중괄호로 열고 닫음.

* 중괄호 내부에는 속성 : 값의 쌍 다수로 선언

  ```css
  h1 { # 선택자. 스타일을 지정할 html 요소를 선택
    color: blue; # 선언 : 속성, 값
    font-size: 15px;
  }
  ```

## CSS 정의 방법

1. 인라인 : 해당 태그에 직접 style 작성
2. 내부 참조 : head 태그 내에 style에 지정 (임베딩 방식)
3. 외부 참조 : 외부 css 파일을 \<head> 내 \<link>를 통해 불러오기

## 내부 참조와 외부 참조

* 외부 참조의 장점 : 내부 참조에서의 중복 문제를 줄일 수 있음
* 그러나 사용 상황에 따라 세개 중 필요한 방식이 달라짐
* 외부 참조를 많이 쓰게 됨

## 속성 이름들

* 속성들 : 엄청 많음. 전부 외워야 하나?
  * 주로 활용하는 속성의 수는 한정되어 있음
  * 빈도가 낮은 속성들은 필요할 때 검섹헤서 사용하면 됨

# CSS Selector

## 선택자

* html 문서에서 특정한 요소를 선택하여 스타일링 하기 위해서는 반드시 선택자라는 개념이 필요하다.
* 기본 선택자
  * 전체 선택자, 요소 선택자
  * 클래스 선택자, 아이디 선택자, 속성 선택자
* 결합자 (Combinators)
  * 자손 결합자, 자식 결합자
  * 일반 형제 결함자, 인접 형제 결합자
* 의사 클래스/요소(pseudo class)
  * 링크, 동적 의사 클래스
  * 구조적 의사 클래스

## 선택자 with 개발자 도구

* chrome에서 활용
* 개발 환경 : VSC

## 선택자들

* \* :  전체 선택자

  ```css
  <style>
  /* 전체 선택자 */
  * {
    color: red'
  }
  
  /* 요소 선택자*/
  h2 {
    color: orange;
  }
  ```

* 각각의 클래스와 아이디에 맞춘 스타일을 작성하도록 도와줌

  ```html
  <style>
  /* 클래스 선택자 */
    .green {
        color: green;
      }
  /* 아이디 선택자 */
    #purple {
        color: purple;
    }
  </style>
  ```

  * .green(예시) : . 이 붙으면 클래스 선택자
    * ex : li.frst : first라는 클래스
  * \# : id를 표현
  * id는 원칙상 선택자 당 한 번만 사용해야하나 동작은 함
  * 클래스는 여러번 사용할 수 있도록 함

  * 복수의 클래스는 공백으로 나열 가능

    ```css
    /* green과 box각각의 클래스를 가져옴*/
    <div class="green box">
    </div>

### CSS 선택자 정리

* 요소 선택자
  * html 태그를 직접 선택
* 클래스(class)선택자
  * 마침표(.) 문자로 시작하며, 해당 클래스가 적용된 모든 항목을 선택
* 아이디(id) 선택자
  * \# 문자로 시작하며, 해당 아이디가 적용된 모든 항목을 선택
  * 일반적으로 하나의 문서에 1번만 사용(Unique)
  * 여러 번 사용해도 동작하지만, 단일의 의미로 쓰는 것이기 때문에 단일 id 사용 권장

## CSS 적용 우선 순위 (cascading order)

* css 우선 순위를 아래와 같이 그룹을 지어볼 수 있다.
  1. 중요도(importance) - 반드시 필요한 경우가 아니면 잘 쓰지 않는다.
     * !important 로 표기
     * cascading order를 깨는 표기법
  2. 우선순위 (Specificity)
     * 인라인 > id 선택자 > class 선택자 > tag 선택자
  3. 거의 class가 많이 사용됨

## CSS 상속

* CSS는 상속을 통해 부모 요소의 속성을 자식에게 상속한다.
  * 속성 중에는 상속이 되는 것과 되지 않는 것들이 있다
  * 상속되는 것 예시
    * 텍스트 관련 요소( font, color, text-align), opcity, visibility 등
  * 상속되지 않는 것 예시
    * 보여지는 것과 관련 요소
      * Box model 관련요소 (width, height, margin, padding, border, box-sizing, display ...)
      * Position 관련 요소

# CSS 단위

## 크기 단위

* px(픽셀)

  * 모니터 해상도의 한 화소인 픽셀을 기준
  * 픽셀의 크기는 변하지 않기 때문에 고정적인 단위
    (실제 출력 크기가 고정되어 있다는 말은 아님)

* %

  * 백분율 단위
  * 가변적인 레이아웃에서 자주 사용

* em

  * (바로 위 부모 요소에 대한) 상속의 영향을 받음
  * 배수 단위, 요소에 지정된 사이즈에 상대적인 사이즈를 가짐.
  * 2em => 부모 10px 일경우 20px

* rem

  * 최상위 부모 요소에 대한 상속 영향을 받음
  * 최상위 요소(html)의 사이즈(16px)를 기준으로 배수 단위의 크기를 가짐

* viewport

  * 웹 페이지를 방문한 유저에게 바로 보이게 되는 웹 컨텐츠의 영역
  * 주로 스마트폰이나 태블릿 디바이스의 화면을 일컫는 용어로 사용됨
  * 글자 그대로 디바이스의 viewport의 사이즈에 따라 상대적인 크기가 결정됨
  * vw, vh, vmin, vmax

  ```html
  <html>
    <head>
      <style>
          .em {
              font-size: 1.5em;
          }
          
          .rem {
              font-size: 1.5rem;
          }
      </style>    
    </head>
    <body>
      <ul class="em">
          <li class="em">1.5em</li>
          <li class="rem">1.5rem</li>
      </ul>    
    </body>
  </html>
  ```

  * 1.5em 의 사이즈는 36, 1.5rem의 사이즈는 24
  * html 디폴트 사이즈는 16

## 색상 단위

1. 색상키워드
   * 대소문자를 구분하지 않음
   * red, blue, black과 같은 특정색 직접 글자로 나타냄
2. RGB 색상
   * 16진수 표기법 혹은 함수형 표기법을 사용해서 특정 색을 표현하는 방식
3. HSL 색상
   * 색상, 채도, 명도를 통해 특정 색을 표현하는 방식

* a 는 alpha(투명도)가 추가된 것

  ```html
  <style>
  p { color : black }
  p { color : #000000 }
  p { color: rgb(0, 0, 0); }
  </style>
  ```

## CSS 문서표현

* 텍스트
  * 변형 서체

# 선택자 심화

## 결합자 (Cominators)

* 자손 결합자

  * selectorA 하위의 모든 selectorB 요소

    ```html
    <body>
       div span {
        	color : red;
    	}
    </body>
    ```

* 자식 결합자

  * selectorA 바로 아래의 selectorB 요소

    ```html
    <body>
        div > span {
    		color : blue;
    	}
    </body>
    ```

* 일반 형제 결합자

  * selectorA의 형제 요소 중 뒤에 위치하는 selectorB 요소를 모두 선택

    ```html
    <body>
        p ~ span {
    		color : blue;
    	}
    </body>
    
    ```

* 인접 형제 결합자

  * selectorA의 형제 요소 중 바로 뒤에 위치하는 selectorB 요소를 선택

    ```html
    <body>
        p + span {
    		color : blue;
    	}
    </body>
    ```

  # CSS Box model

  ## 사각형의 세계

  * 모든 HTML 요소는 box 형태로 되어있음

  * 하나의 박스는 네 부분으로 이루어짐

  * margin

    * 테두리 바깥의 외부 **여백**
    * 배경색을 지정할 수 없다

  * border

    * 테두리 영역

  * padding

    * 테두리 안쪽의 내부 여백 요소에 적용된 배경색, 이미지는 padding 이미지 적용
    * content 부터 테두리까지의 거리

  * content

    * 글이나 이미지 등 요소의 실제 내용

    ```css
    <style>
    .margin {
      margin-top: 10px;
      margin-right: 20px;
      margin-bottom: 30px;
      margin-left: 40px;
    }
    
    .margin-padding {
      /* 별도로 지정하지 않으면 상하좌우 균일하게 10px 적용*/
      margin: 10px
      padding: 30px
    }
    
    /* short hand */
    .margin-2 {
      margin: 10px 20px;
      /*상하/좌우*/
    }
    .margin-3 {
      margin:10px 20px 30px;
      /*상/좌우/하*/
    }
    .margin-4 {
      margin: 10px 20px 30px 40px;
      /*상/우/좌/하*/
    }
    
    .border {
      border-width: 2px;
      border-style: dashed;
      border-color: black;
    }
    /* == */
    .border {
      border: 2px dashed black;
    }
    
    .box {
      width: 100px;
      margin: 10px auto
      padding: 20px
      border: 1px
    }
    /* content box의 너비는 1 + 20 + 100 + 20 + 1 = 142 */
    
    /* content box q
    .box-sizing {
      box-sizing: border-box
    }
    </style>
    ```



## marin 상쇄

* 인접 형제 요소간의 마진을 겹쳐서 표현
* 일반 인식은 마진 20이면 다른 요소간의 거리가 40이 아닌 20이기 때문



## 인라인 / 블락

* display: block
  * 줄 바꿈이 일어나는 요소



## 속성에 따른 수평 정렬

* margin-right: auto; -> text-align :left;
* margin-left auto; -> text- align right

# CSS Display

* 요소들을 시각적으로 어떻게 보여줄지 결정하는 속성

# CSS Posiion

## CSS Position

* relative : 상대 위치
  * 자기 자신의 static 위치를 기준으로 이동
  * 레이 아웃에서 요소가 차지 하는 공간은 static일 때와 같음
* absolute: 절대 위치
  * 요소를 일반적인 문서 흐름에서 제거 후 레이아웃에 공간을 차지하지 않음
  * static이 아닌 가장 가까운 부모 / 조상 요소를 기준으로 이동
* fixed : 고정 위치
  * 요소를 일반적인 문서 흐름ㅇ레서 제거 후 레이아웃에 공간을 차지 하지 않음
  * 부모요소와 관계 없이 viewprot를 기준으로 이동
  * 스크롤 시에도 항상 같은 곳에 위치함

## Absolute vs Relative

* 