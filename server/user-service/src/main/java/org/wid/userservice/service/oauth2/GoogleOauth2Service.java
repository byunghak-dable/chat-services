package org.wid.userservice.service.oauth2;

import org.springframework.beans.factory.annotation.Qualifier;
import org.springframework.stereotype.Service;
import org.wid.userservice.dto.user.OauthLoginResponseDto;

import lombok.extern.slf4j.Slf4j;
import reactor.core.publisher.Mono;

@Service
@Qualifier("GoogleOauth2Service")
@Slf4j
public class GoogleOauth2Service implements Oauth2Service {

  @Override
  public Mono<OauthLoginResponseDto> getToken(String code) {
    log.info("google code: {}", code);
    throw new UnsupportedOperationException("Unimplemented method 'getToken' in google");
  }
}
