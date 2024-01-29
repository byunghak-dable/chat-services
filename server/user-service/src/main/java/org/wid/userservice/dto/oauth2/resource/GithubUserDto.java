package org.wid.userservice.dto.oauth2.resource;

public record GithubUserDto(
    int id,
    String login,
    String nodeId,
    String avatarUrl,
    String gravatarId,
    String url,
    String htmlUrl,
    String followersUrl,
    String followingUrl,
    String gistsUrl,
    String starredUrl,
    String subscriptionsUrl,
    String organizationsUrl,
    String reposUrl,
    String eventsUrl,
    String receivedEventsUrl,
    String type,
    Boolean siteAdmin,
    String name,
    String company,
    String blog,
    String location,
    String email,
    String hireable,
    String bio,
    String twitterUsername,
    int publicRepos,
    int publicGists,
    int followers,
    int following,
    String createdAt,
    String updatedAt

) {
}
