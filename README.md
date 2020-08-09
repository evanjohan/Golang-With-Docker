# Golang-With-DockerFile
  Mengupload Images App Golang ke Docker Registry (Docker Hub).
  Dengan Cara:<br/>
    1. Menggunakan DockerFile yang tersedia, dengan syntax :<br/>
    &nbsp;&nbsp;&nbsp;&nbsp;> Docker build --tag [name Image : tagname] .  <br/>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    Example : app-golang:1.0 <br/>
    2. Upload ke docker registry, dengan syntax :<br/>
    &nbsp;&nbsp;&nbsp;&nbsp;> Docker  docker push [reponame : tagname] <br/>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    Example : evanjohan/app-golang:1.0<br/>
    3. Jika akses denied, maka gunakan syntax: <br/>
    &nbsp;&nbsp;&nbsp;&nbsp;> Docker login (nanti masukan username dan password docker hub).

<h>Keterangan:<h/><br/> 
    >    Nama image harus sama dengan nama repository yang ada di docker hub. Contoh nama image **evanjohan/app-golang** , maka repositorynya juga **evanjohan/app-golang**.
