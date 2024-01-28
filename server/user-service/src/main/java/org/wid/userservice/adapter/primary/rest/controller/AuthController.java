package org.wid.userservice.adapter.primary.rest.controller;

import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.wid.userservice.dto.auth.Oauth2LoginRequestDto;
import org.wid.userservice.port.primary.AuthServicePort;

import jakarta.validation.Valid;
import lombok.RequiredArgsConstructor;
import reactor.core.publisher.Mono;

@RestController
@RequestMapping("/auth/v1")
@RequiredArgsConstructor
public class AuthController {
  private final AuthServicePort authService;

  @PostMapping("/login/oauth2")
  public Mono<Object> googleLogin(@Valid @RequestBody Oauth2LoginRequestDto loginRequestDto)
      throws IllegalArgumentException {

    return authService.oauth2Login(loginRequestDto);
  }
}
