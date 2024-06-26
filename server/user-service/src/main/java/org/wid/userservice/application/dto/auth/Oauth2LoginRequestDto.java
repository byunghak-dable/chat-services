package org.wid.userservice.application.dto.auth;

import jakarta.validation.constraints.NotEmpty;
import jakarta.validation.constraints.NotNull;
import lombok.Getter;
import org.wid.userservice.domain.entity.User.LoginType;

@Getter
public class Oauth2LoginRequestDto {
  @NotNull private final LoginType type;

  @NotEmpty private final String code;

  public Oauth2LoginRequestDto(String type, String code) {
    this.type = LoginType.valueOf(type.toUpperCase());
    this.code = code;
  }
}
