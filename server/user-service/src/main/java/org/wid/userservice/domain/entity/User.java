package org.wid.userservice.domain.entity;

import lombok.Builder;
import lombok.Getter;
import lombok.RequiredArgsConstructor;

@Getter
@Builder
public record User(String id, String email, String name, String profile, LoginType loginType) {
  public enum LoginType {
    GOOGLE,
    GITHUB
  }
}
