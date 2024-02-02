package org.wid.userservice.port.primary;

import org.wid.userservice.dto.user.UserDto;

public interface JwtServicePort {
  String createAccessToken(UserDto userDto);

  String createRefreshToken(UserDto userDto);
}
