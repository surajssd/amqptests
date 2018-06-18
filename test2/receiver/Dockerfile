FROM fedora:28

RUN dnf -y install qpid-proton-c-devel && dnf clean all

COPY receiver /receiver

CMD /receiver
