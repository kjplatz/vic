# Supported Docker Commands

 vSphere Integrated Containers supports Docker 1.11.2. The supported version of the Docker API is 1.23. If you are using a more recent version of the Docker client, see [Docker Commands Fail with a Docker API Version Error](ts_docker_version_error.md).

## General Container Operations ##


| **Container** | **Docker Reference** | **Supported** |
| --- | --- | --- |
|Docker start|[Start a container](https://docs.docker.com/engine/reference/commandline/start/)|Yes|
|Docker version|[Docker version information](https://docs.docker.com/engine/reference/commandline/version/)|Yes. vSphere Integrated Containers version provided|
|Docker run|Composite command of create, start, inspect, attach, rm, resize, wait, kill|Yes. <code>docker run -c</code> and <code>docker run -m</code> parameters are supported|
|Docker create|[Create a container](https://docs.docker.com/engine/reference/commandline/create/)|Yes|
|Docker inspect|[Inspect a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#inspect-a-container) <br>[Inspect an image](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#inspect-an-image)|Yes|
|Docker port|[Obtain port data](https://docs.docker.com/engine/reference/commandline/port/)|Yes. Displays port mapping data.|
|Docker attach|[Attach to a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#attach-to-a-container)<br> [Attach to a container websocket](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#attach-to-a-container-websocket)|Yes|
|Docker rm|[Remove a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#remove-a-container)|Yes, only the <code>name</code> parameter is supported. <code>force</code> and <code>v</code> are a future implementation. Also removes associated volumes.|
|Docker info|[Docker system information](https://docs.docker.com/engine/reference/commandline/info/)|Yes, docker-specific data, basic capacity information, list of configured volume stores, virtual container host information. Does not reveal vSphere datastore paths that might contain sensitive vSphere information|
|Docker container resize|[Resize a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.23/#resize-a-container-tty)|Yes|
|Docker stop|[Stop a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#stop-a-container)<br> [Stop](https://docs.docker.com/engine/reference/commandline/stop/)|Yes. Powers down the VM |
|Docker images|[Images](https://docs.docker.com/engine/reference/commandline/images/)<br>[list-images](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#list-images)|Yes|
|Docker ps|[Show running containers](https://docs.docker.com/engine/reference/commandline/ps/)|Yes|
|Docker logs|[Get container logs](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#get-container-logs)|Yes, except for the `docker logs --timestamps` (`-t`) and `--since` options, which are not supported.|
|Docker restart|[Restart a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#restart-a-container)<br> [Restart](https://docs.docker.com/engine/reference/commandline/restart/)|Yes|
|Docker kill|[Kill a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#kill-a-container)<br> [Kill](https://docs.docker.com/engine/reference/commandline/kill/)|No. Use Docker stop|
|Docker Container list|[List Containers](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#list-containers)|Yes|

## Network Operations ##

For more information about network operations, see [Network Port Use Cases](network_use_cases.md).

| **Network** | **Docker Reference** | **Supported** |
| --- | --- | --- |
|Network create|[Create a network](https://docs.docker.com/engine/reference/commandline/network_create/)|Yes. See the use case to connect to an external network in vSphere Integrated Container for vSphere Administrators. Bridge is also supported.|
|Network ls|[List networks/](https://docs.docker.com/engine/reference/commandline/network_ls/)|Yes|
|Network inspect|[Inspect a network](https://docs.docker.com/engine/reference/commandline/network_inspect/)|Yes|
|Network connect|[Connect to a network](https://docs.docker.com/engine/reference/commandline/network_connect/)|Yes|
|Network rm|[Remove a network](https://docs.docker.com/engine/reference/commandline/network_rm/)|Yes|

## Volume Operations ##

For more information about volume operations, see [Using Volumes with vSphere Integrated Containers](using_volumes_with_vic.md).

| **Volume** | **Docker Reference** | **Supported** |
| --- | --- | --- |
|Docker volume create|[Create a volume](https://docs.docker.com/engine/reference/commandline/volume_create/)| The driver option is ignored even if you specify it. You must include <code>--opt VolumeStore=<name of volumestore> --Capacity=<size  in MB></code> as these are direct vSphere arguments. VIC does not assign random names during a volume create, but only for anonymous volumes.|
|Docker volume ls|[List volumes](https://docs.docker.com/engine/reference/commandline/volume_ls/) | Yes |
|Docker volume rm|[Remove a volume](https://docs.docker.com/engine/reference/commandline/volume_rm/) | No |

## Other Operations ##

| **Commands** | **Docker Reference** | **Supported** |
| --- | --- | --- |
|Link|[Link](https://docs.docker.com/v1.8/userguide/dockerlinks/)|Future release|
|Docker export|[Export a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#export-a-container)|Future release|
|Docker save|[Save images](https://docs.docker.com/engine/reference/commandline/save/)|Future release|
|Docker stats|[Get container stats based on resource usage](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#get-container-stats-based-on-resource-usage)<br> [Stats](https://docs.docker.com/engine/reference/commandline/stats/)|Future release|
|Docker update| [Update a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#update-a-container) <br> [Update](https://docs.docker.com/engine/reference/commandline/update/)|Future release|
|Docker rename|[Rename a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#rename-a-container) [Rename](https://docs.docker.com/engine/reference/commandline/rename/)|Future release|
|Docker pause|[Pause processes in a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#pause-a-container)<br> [Pause](https://docs.docker.com/engine/reference/commandline/pause/)|Future release|
|Docker unpause|[Unpause processes in a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#unpause-a-container)<br> [Unpause](https://docs.docker.com/engine/reference/commandline/unpause/)|Future release|
|Docker cp| [Copy files or folders in a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#copy-files-or-folders-from-a-container) <br> [Copy](https://docs.docker.com/engine/reference/commandline/cp/)|Future release|
|Wait|[Wait for a container](https://docs.docker.com/engine/reference/api/docker_remote_api_v1.22/#wait-a-container) <br>[Wait](https://docs.docker.com/engine/reference/commandline/wait/)|Future release|