import React from 'react';
import { BottomNavigation } from 'react-native-paper';
import Friends from './Friends';
import Match from './Match';
import Profile from './Profile';

export default function Home() {
  const [index, setIndex] = React.useState(0);
  const [routes] = React.useState([
    { key: 'match', title: 'Match', focusedIcon: 'heart', unfocusedIcon: 'heart-outline' },
    { key: 'friends', title: 'Friends', focusedIcon: 'friends' },
    { key: 'profile', title: 'Profile', focusedIcon: 'profile' },
  ]);

  return (
    <BottomNavigation
      navigationState={{ index, routes }}
      onIndexChange={setIndex}
      renderScene={BottomNavigation.SceneMap({
        match: Match,
        friends: Friends,
        profile: Profile,
      })}
    />
  );
}
