package org.wid.userservice.service;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.entity.entity.User.LoginType;
import org.wid.userservice.port.primary.AuthServicePort;
import org.wid.userservice.service.oauth2.Oauth2Service;

import reactor.core.publisher.Mono;

@Service
public class AuthService implements AuthServicePort {

  private final Map<LoginType, Oauth2Service> oauth2ServiceMap;

  public AuthService(@Qualifier("GoogleOauth2Service") Oauth2Service googleOauth2Service) {
    oauth2ServiceMap = Map.of(LoginType.GOOGLE, googleOauth2Service);
  }

  @Override
  public Mono<Void> oauth2Login(Oauth2LoginRequestDto loginDto) {
    Oauth2Service oauth2Service = oauth2ServiceMap.get(loginDto.getType());

    oauth2Service.requestAccessToken(loginDto.getCode());
    return Mono.empty();
  }
}
