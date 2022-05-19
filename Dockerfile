FROM ubuntu:18.10 AS base

COPY bin/app/cmd /cmd

CMD ["/cmd"]
