package org.wid.userservice.port.driving;

import org.wid.userservice.application.dto.auth.JwtDto;
import org.wid.userservice.application.dto.user.UserDto;

public interface JwtServicePort {
  JwtDto generateTokens(UserDto userDto);

  JwtDto refresh(String refreshToken);
}
