package org.wid.userservice.entity.entity;

import jakarta.persistence.Entity;

public class User {
  int id;
  String email;
  String password;
  String name;
  Boolean gender;
  String thumbnailUrl;
}
