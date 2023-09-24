# sha256:25be7e009d19517822a7bebc361c0c7c88f7a087396056760d6665058ba8440e
# "org.opencontainers.image.revision": "76376c9148ed65b0195be8033ed7bc2ada5a6230",
# "org.opencontainers.image.source": "https://github.com/zmkfirmware/zmk-docker"

FROM zmkfirmware/zmk-build-arm:3.2-branch

WORKDIR /west

COPY config/west.yml ./config/west.yml
RUN west init -l "./config"

RUN west update --narrow
RUN west zephyr-export

COPY tools/build.sh /usr/bin/build.sh
ENTRYPOINT [ "/usr/bin/build.sh" ]

