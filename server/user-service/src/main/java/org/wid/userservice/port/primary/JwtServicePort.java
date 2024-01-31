package org.wid.userservice.port.primary;

public interface JwtServicePort {
  String createAccessToken();

  String createRefreshToken();
}
