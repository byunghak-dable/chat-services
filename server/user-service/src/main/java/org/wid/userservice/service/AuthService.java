package org.wid.userservice.service;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.OauthLoginResponseDto;
import org.wid.userservice.port.primary.AuthServicePort;
import org.wid.userservice.service.oauth2.Oauth2Service;

import reactor.core.publisher.Mono;

@Service
public class AuthService implements AuthServicePort {

  private final Oauth2Service googleService;
  private final Oauth2Service githubService;

  public AuthService(
      @Qualifier("GoogleOauth2Service") Oauth2Service googleService,
      @Qualifier("GithubOauth2Service") Oauth2Service githubService) {
    this.googleService = googleService;
    this.githubService = githubService;
  }

  @Override
  public Mono<OauthLoginResponseDto> googleLogin(String code) {
    return googleService.getToken(code);
  }

  @Override
  public Mono<OauthLoginResponseDto> githubLogin(String code) {
    return githubService.getToken(code);
  }
}
