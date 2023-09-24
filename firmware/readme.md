

```
docker run --rm -it \
  -v $PWD/config:/config \
  -v $PWD/bin:/west/bin \
  -v $PWD/cache:/root/.cache/zephyr \
  zmkbuilder:local
```