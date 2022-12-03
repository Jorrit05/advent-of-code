#!/bin/bash

[[ -z $1 ]] && { echo "usage: createRascalProject <projectName>"; exit 1; }
name=$1

[[ -f $name || -d $name ]] && { echo "$name already exists"; exit 1; }
mkdir -p "$name"
touch "${name}/pom.xml"
mkdir -p "${name}/src/main/rascal"
touch "${name}/src/main/rascal/Main.rsc"
mkdir -p "${name}/META-INF"
touch "${name}/META-INF/RASCAL.MF"

cat > "${name}/META-INF/RASCAL.MF" << EOM
Manifest-Version: 0.0.1
Project-Name: ${name}
Source: src/main/rascal
Require-Libraries: 
EOM

cat > "${name}/src/main/rascal/Main.rsc" << EOM
module Main

import IO;

void main() {
    println("Hello world");
}
EOM

cat > "${name}/pom.xml" << EOM
<?xml version="1.0" encoding="UTF-8"?>
  <project xmlns="http://maven.apache.org/POM/4.0.0" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xsi:schemaLocation="http://maven.apache.org/POM/4.0.0 http://maven.apache.org/xsd/maven-4.0.0.xsd">
  <modelVersion>4.0.0</modelVersion>

  <groupId>org.rascalmpl</groupId>
  <artifactId>${name}</artifactId>
  <version>0.1.0-SNAPSHOT</version>

  <properties>
    <project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>
  </properties>

  <repositories>
    <repository>
        <id>usethesource</id>
        <url>https://releases.usethesource.io/maven/</url>
    </repository>
  </repositories>

  <pluginRepositories>
    <pluginRepository>
       <id>usethesource</id>
       <url>https://releases.usethesource.io/maven/</url>
    </pluginRepository>
  </pluginRepositories>

  <dependencies>
    <dependency>
      <groupId>org.rascalmpl</groupId>
      <artifactId>rascal</artifactId>
      <version>0.26.2</version>
    </dependency>
  </dependencies>

  <build>
    <plugins>
      <plugin>
        <groupId>org.apache.maven.plugins</groupId>
        <artifactId>maven-compiler-plugin</artifactId>
        <version>3.8.0</version>
        <configuration>
          <compilerArgument>-parameters</compilerArgument> 
          <release>11</release>
        </configuration>
      </plugin>
      <plugin>
        <groupId>org.rascalmpl</groupId>
        <artifactId>rascal-maven-plugin</artifactId>
        <version>0.8.2</version>
        <configuration>
          <errorsAsWarnings>true</errorsAsWarnings>
          <bin>\${project.build.outputDirectory}</bin>
          <srcs>
            <src>\${project.basedir}/src/main/rascal</src>
          </srcs>
        </configuration>
      </plugin>
    </plugins>
  </build>
</project>
EOM