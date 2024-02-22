package org.wid.userservice.application.service;

import java.util.Map;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.application.dto.auth.AuthenticationTokensDto;
import org.wid.userservice.application.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.application.service.oauth2.Oauth2Service;
import org.wid.userservice.domain.entity.User.LoginType;
import org.wid.userservice.port.driving.AuthServicePort;
import org.wid.userservice.port.driving.UserServicePort;

import reactor.core.publisher.Mono;

@Service
public class AuthService implements AuthServicePort {

  private final UserServicePort userService;
  private final TokenService jwtService;
  private final Map<LoginType, Oauth2Service> oauth2ServiceMap;

  public AuthService(
      UserServicePort userService,
      TokenService jwtService,
      @Qualifier("GoogleOauth2Service") Oauth2Service googleOauth2Service,
      @Qualifier("GithubOauth2Service") Oauth2Service githubOauth2Service) {
    this.userService = userService;
    this.jwtService = jwtService;
    this.oauth2ServiceMap = Map.of(
        LoginType.GOOGLE, googleOauth2Service,
        LoginType.GITHUB, githubOauth2Service);
  }

  @Override
  public Mono<AuthenticationTokensDto> oauth2Login(Oauth2LoginRequestDto loginDto) {
    Oauth2Service oauth2Service = oauth2ServiceMap.get(loginDto.getType());

    return oauth2Service
        .getToken(loginDto.getCode())
        .flatMap(oauth2Service::getResource)
        .flatMap(userService::upsertUser)
        .map(jwtService::generateTokens);
  }

  @Override
  public AuthenticationTokensDto generateAccessToken(String refreshToken) {
    return jwtService.generateAccessToken(refreshToken);
  }
}
