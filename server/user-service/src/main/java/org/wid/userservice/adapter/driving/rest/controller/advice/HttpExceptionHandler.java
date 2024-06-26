package org.wid.userservice.adapter.driving.rest.controller.advice;

import org.springframework.http.HttpStatus;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseStatus;
import org.springframework.web.bind.annotation.RestControllerAdvice;
import org.wid.userservice.adapter.driving.rest.controller.AuthController;
import org.wid.userservice.adapter.driving.rest.controller.UserController;
import org.wid.userservice.application.dto.ErrorResponseDto;
import org.wid.userservice.application.exception.BadRequestException;
import reactor.core.publisher.Mono;

@RestControllerAdvice(assignableTypes = {UserController.class, AuthController.class})
public class HttpExceptionHandler {

  @ExceptionHandler(Exception.class)
  @ResponseStatus(HttpStatus.INTERNAL_SERVER_ERROR)
  public Mono<ErrorResponseDto> handlDefaultException(Exception e) {
    return Mono.just(new ErrorResponseDto(HttpStatus.INTERNAL_SERVER_ERROR, e.getMessage()));
  }

  @ExceptionHandler({BadRequestException.class, IllegalArgumentException.class})
  @ResponseStatus(HttpStatus.BAD_REQUEST)
  public Mono<ErrorResponseDto> handleBadRequest(Exception e) {
    return Mono.just(new ErrorResponseDto(HttpStatus.BAD_REQUEST, e.getMessage()));
  }
}
