FROM moby/buildkit:v0.9.3
WORKDIR /errors
COPY errors README.md /errors/
ENV PATH=/errors:$PATH
ENTRYPOINT [ "/bhojpur/errors" ]