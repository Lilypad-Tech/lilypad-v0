FROM ubuntu
RUN apt update && apt install -y jq bc && rm -rf /var/lib/apt/lists/*
COPY rand /bin/rand
CMD /bin/rand
