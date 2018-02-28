# 2018-02-27 (Golang/Modern Devops) Paris meetup

Files for the 2018-02-27 Docker talk of the meta
[(Golang/Modern Devops) Paris meetup](https://www.meetup.com/fr-FR/Meetup-Modern-Devops-Paris/events/247749767/).

## Slides

Here is the infamouse Prezi presentation :
https://prezi.com/p/2-kdk2x0pnfy/

:warning: Sickness bags not provided.

## Steps

1. A standard docker image with the default Golang image
2. Reordered Dockerfile to benefit cache
3. The `onbuild` special image
4. Choose the file to add or use .dockerignore
5. The `golang:alpine` image
6. Only add the binary to the image
7. Use docker multi-stage to build the binary and keep it in the final image
8. Use docker multi-stage to build multiple versions of an image  
9. Create your own base image to avoid RUN steps (not cached)

## Build & run

For each step, you can build the image with :

```bash
docker build -t <image name> -f <step directory>/Dockerfile .
```

For example :

```bash
docker build -t meetup3 -f 3.onbuild/Dockerfile .
```

Once built, you can run the image :
```bash
docker run --rm -ti -p 8080:8080  meetup6
# In another terminal
curl -si http://127.0.0.1:8080
```

For the step 8 :

```bash
# Build the production image
docker build --target=production -t meetup8:production -f 8.multiple_versions/Dockerfile .
# Build the debug image
docker build --target=debug -t meetup8:debug -f 8.multiple_versions/Dockerfile .
```

The debug image must be launch with :

```bash
docker run --rm -ti -p 8080:8080 -p 2345:2345 --security-opt=seccomp:unconfined meetup8:debug
```

Then you can [remote debug](https://github.com/Microsoft/vscode-go/wiki/Debugging-Go-code-using-VS-Code#remote-debugging) with VSCode !