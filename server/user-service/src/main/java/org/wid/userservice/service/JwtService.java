package org.wid.userservice.service;

import org.springframework.stereotype.Service;
import org.wid.userservice.port.primary.JwtServicePort;

@Service
public class JwtService implements JwtServicePort {

  @Override
  public String createAccessToken() {
    // TODO Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'createAccessToken'");
  }

  @Override
  public String createRefreshToken() {
    // TODO Auto-generated method stub
    throw new UnsupportedOperationException("Unimplemented method 'createRefreshToken'");
  }
}
