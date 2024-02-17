package org.wid.userservice.entity;

import org.springframework.data.annotation.Id;
import org.springframework.data.mongodb.core.mapping.Document;

import lombok.Builder;
import lombok.Getter;

@Document(collection = "user")
@Getter
@Builder
public class User {

  public enum LoginType {
    GOOGLE, GITHUB
  }

  @Id
  private String id;

  private String email;

  private String name;

  private String profile;

  private LoginType loginType;
}
