@echo off
REM 启动第一个节点
start "" go run ..\node\main.go --port=8001
timeout /t 2 /nobreak

REM 启动第二个节点
start "" go run ..\node\main.go --port=8002
timeout /t 2 /nobreak

REM 启动第三个节点
start "" go run ..\node\main.go --port=8003
timeout /t 2 /nobreak

REM 启动第四个节点
start "" go run ..\node\main.go --port=8004
timeout /t 2 /nobreak

REM 启动第五个节点
start "" go run ..\node\main.go --port=8005
timeout /t 2 /nobreak
