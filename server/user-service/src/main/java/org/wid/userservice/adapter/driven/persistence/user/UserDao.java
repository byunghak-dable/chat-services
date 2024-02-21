package org.wid.userservice.adapter.driven.persistence.user;

import org.springframework.data.mongodb.repository.ReactiveMongoRepository;
import org.wid.userservice.domain.entity.User;

interface UserDao extends ReactiveMongoRepository<User, String> {
}
