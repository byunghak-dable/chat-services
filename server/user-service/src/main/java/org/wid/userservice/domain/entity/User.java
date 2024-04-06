package org.wid.userservice.domain.entity;

public record User(String id, String email, String name, String profile, LoginType loginType) {
  public enum LoginType {
    GOOGLE,
    GITHUB
  }
}
