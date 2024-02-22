package org.wid.userservice.application.abstraction;

import org.wid.userservice.application.dto.auth.AccessTokenDto;
import org.wid.userservice.application.dto.auth.AuthenticationTokensDto;
import org.wid.userservice.application.dto.user.UserDto;

public interface TokenService {
  AuthenticationTokensDto generateTokens(UserDto userDto);

  AccessTokenDto generateAccessToken(String refreshToken);
}
