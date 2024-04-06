package org.wid.userservice.domain.entity;

import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@RequiredArgsConstructor
public class User {
  private final String id;
  private final String email;
  private final String name;
  private final String profile;
  private final LoginType loginType;

  public enum LoginType {
    GOOGLE,
    GITHUB
  }
}
