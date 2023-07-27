# Modbus Protocol (2023/07/27)  

Modbus Protocol에 대해 정리합니다.  

- [Modbus Protocol (2023/07/27)](#modbus-protocol-20230727)
  - [Setup](#setup)
  - [What is Modbus?](#what-is-modbus)
  - [Memory Map](#memory-map)
  - [Modbus Packet Frame Format](#modbus-packet-frame-format)
  - [Get Started](#get-started)
  - [Wireshark Packet Analyze](#wireshark-packet-analyze)

## Setup
- Wireshark
  - [Downlaod Site](https://www.wireshark.org/download.html)
- Modbus Poll v10.5.0
  - [Downlaod Site](https://www.modbustools.com/download.html)
- Modbus Slave v8.2.2
  - [Downlaod Site](https://www.modbustools.com/download.html)

## What is Modbus?
Modbus는 산업용 통신 프로토콜입니다.  
마스터(이하 클라이언트)-슬레이브(이하 서버) 계층의 구조를 가집니다.
<br>

## Memory Map
Modbus의 가장 큰 특징 중 하나인 Memory Map입니다.  
Memory Map은 클라이언트와 서버가 공유하며, 데이터를 읽기/쓰기 하는 데에 사용됩니다.  
Memory Map Convention은 장비에 따라 다를 수 있겠지만, 모두 Modbus Protocol을 따릅니다.  
이에 따라 Modbus를 지원하는 모든 장비는, Protocol을 별도로 수정하지 않아도 됩니다.  

|함수 번호|접근     |타입  |용도             |코일/레지스터 메모리 구역            |
|--:|:-------------|:----:|:---------------|:----------------------------------|
|01 |Read/Write    |1 bit |Read Coil               |00001 ~ 09999              |
|05 |Read/Write    |1 bit |Write Single Coil       |00001 ~ 09999              |
|15 |Read/Write    |1 bit |Write Multiple Coils    |00001 ~ 09999              |
|02 |Read Only     |1  bit|Read Discrete Inputs    |10001 ~ 19999              |
|03 |Read/Write    |16 bit|Read Holding Register   |40001 ~ 49999              |
|06 |Read/Write    |16 bit|Write Single Register   |40001 ~ 49999              |
|16 |Read/Write    |16 bit|Write Multiple Register |40001 ~ 49999              |
|04 |Read Only     |16 bit|Read Input Register     |30001 ~ 39999              |

코일은 1bit의 데이터(예를 들어 LED On/Off Status) 입출력에 사용됩니다.  
레지스터는 16bit(1 Word)의 데이터 입출력에 사용됩니다.  

Modbus에서는 함수 번호별로 접근하고자 하는 메모리 구역이 분리됩니다.  
따라서 접근하고자 하는 메모리의 구역을 나타낼 때에는 앞 자리를 생략하고, 뒤 4자리를 사용합니다.

> ### 예시    
> ---
> 요구사항
> - 30001번 레지스터의 값을 읽어와야 한다. (3xxxx번 구역에서는 쓰기 연산 불가)   
> 
> 동작과정
>   - 클라이언트가 서버에게 04번 함수와 0번째 메모리 구역을 의미하는 패킷을 전송한다.
>   - 서버는 04번 함수가 접근하는 메모리 구역 중, 0번 째 주소인 30001번째 레지스터에 접근한다.
>   - 데이터를 읽어 클라이언트에게 전송한다. (04번 함수 = 읽기 수행)

<br>

## Modbus Packet Frame Format
![MODBUS-RTU-and-ASCII-frame](https://github.com/goburrow/modbus/assets/53927414/c73fba5f-c62a-442e-b697-da49eab7e3ae)
[ResearchGate](https://www.researchgate.net/figure/MODBUS-RTU-and-ASCII-frame_fig1_362987592)

> <strong>Modbus의 패킷은 요구되는 함수에 따라 길이가 가변적입니다.</strong>   

Start와 End는 Frame 간 공백을 의미합니다.

## Get Started
![Wireshark-Init](https://github.com/goburrow/modbus/assets/53927414/b76ac5ca-1580-4d2a-afaa-fecc5cd9535c)

Wireshark 패킷 분석기 실행 후 <code>tcp.port == 502</code> 로 필터링합니다.  
<b>(Modbus의 Default Port는 502입니다.)</b>

![Modbus-Poll-Init](https://github.com/goburrow/modbus/assets/53927414/1b6ab6dc-6d7b-4550-92e9-5c5efa87c837)  
Modbus Poll 실행환경을 설정합니다.

![Modbus-Slave-Init](https://github.com/goburrow/modbus/assets/53927414/21c4b3da-2d6f-4edf-ad63-b06846e87dc1)  
Modbus Slave 실행환경을 설정합니다.

![Set-Auto-Increase](https://github.com/goburrow/modbus/assets/53927414/a8b23907-4a33-4043-8f5f-a84485633058)
Modbus Slave의 Read용 Memory Map에 Value 설정 후 Auto Increase를 체크합니다.  
(Name, Map number, Value는 다르게 하셔도 됩니다.)  

![Write](https://github.com/goburrow/modbus/assets/53927414/607c076c-9a65-4602-b3ce-f5577d99857f)  
Modbus Poll의 0번 주소에 Slave로부터 데이터를 Read하고 있는 모습입니다.  
Modbus Poll에서 06번 함수를 이용해 1번 주소에 8080을 Write합니다.  

![In](https://github.com/goburrow/modbus/assets/53927414/6226cb0b-f1e2-490b-aead-327706b6cb16)
Modbus Slave의 1번 주소에서 Poll이 Write한 데이터를 확인할 수 있습니다. 

<br>

## Wireshark Packet Analyze
Wireshar가 포착한 Packet을 분석합니다.  
58055: Master(Client)  
502: Slave(Server)

![3-way-handshake](https://github.com/goburrow/modbus/assets/53927414/3a140e2c-492b-4450-bec5-a81b7d8ea3f4)
3 way handshake

<br>

![wireshark-client-read](https://github.com/goburrow/modbus/assets/53927414/6a1f9372-d604-40dc-94b5-88cb07fdfa84)
클라이언트(58388 포트)에서 서버(502 포트)로 Read를 요청한 패킷을 분석해보겠습니다.  
[RTU Frame 참조](#modbus-packet-frame-format)  

> 01 -> 1번 Slave에게    
03 -> 03번 함수(Register Read only)를 요청하여 (40001 ~ 49999 사이의 Register)  
00 00 -> 0번 주소부터  
00 0a -> a개 주소만큼의 데이터  
c5 cd -> CRC

<br>

![wireshark-server-read](https://github.com/goburrow/modbus/assets/53927414/74ca6bcd-3964-412d-985d-070c9e580900)
서버(502 포트)에서 클라이언트(58388 포트)로 Read에 응답한 패킷을 분석해보겠습니다.  

> 01 -> 1번 Master에게  
03 -> 03번 함수(Register Read only)에 응답  
14 -> 20 Byte의 데이터 전송 (10개 요청 * 2 Byte)  
00 01 -> 0번 주소의 데이터는 1  
00 00 -> 1번 주소의 데이터는 0  
00 00 -> 2번 주소의 데이터는 0  
...  
...  
9e 9b -> CRC

<br>

![wireshark-client-write](https://github.com/goburrow/modbus/assets/53927414/84fb4017-dc72-4f4e-9c20-402ef3d84769)
클라이언트(58388 포트)에서 서버(502 포트)로 Write를 요청한 패킷을 분석해보겠습니다. 

> 01 -> 1번 Slave에게  
06 -> 06번 함수(Write Single Register) 요청하여 (40001 ~ 49999 사이의 Register)  
00 01 -> 1번 주소에 (Memory Map 상에서 2번째 주소)
1f 90 -> 8080 이라는 데이터  
d0 56 -> CRC

<br>

![wireshark-server-write](https://github.com/goburrow/modbus/assets/53927414/4d2b36eb-405b-4f43-810e-21b7878c0592)
서버(502 포트)에서 클라이언트(58388 포트)로 Write에 응답한 패킷을 분석해보겠습니다. 

> 01 -> 1번 Master에게  
06 -> 06번 함수(Write Single Register)에 응답  
00 01 -> 1번 주소에 (Memory Map 상에서 2번째 주소)  
1f 90 -> 8080 이라는 데이터  
d0 56 -> CRC

<br>

![3-way-handshake-finish](https://github.com/goburrow/modbus/assets/53927414/4cfff419-2a77-4d27-bc1e-d98c3b99fb71)  
3 way handshake

<!-- todo
CRC 예외 발생 시 클라이언트 상에서 응답하지 않는 것이 좋음.
- 이미 오염된 패킷임.
- 서버의 과부하 가능성
- 서버 측에서 timeout을 거는 것이 좋아보임 (이것도 효율 생각해야함) -->

