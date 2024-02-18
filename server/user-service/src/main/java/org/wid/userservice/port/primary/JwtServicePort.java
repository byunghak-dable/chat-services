package org.wid.userservice.port.primary;

import org.wid.userservice.dto.auth.JwtDto;
import org.wid.userservice.dto.user.UserDto;

public interface JwtServicePort {
  JwtDto generateTokens(UserDto userDto);

  JwtDto refresh(String refreshToken);
}
