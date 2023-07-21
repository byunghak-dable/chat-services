import { registerRootComponent } from 'expo';
import { SafeAreaProvider } from 'react-native-safe-area-context';
import { NavigationContainer } from '@react-navigation/native';
import { createStackNavigator } from '@react-navigation/stack';
import Signin from './screen/Signin';
import Home from './screen/Home';

const { Navigator, Screen } = createStackNavigator();

export default registerRootComponent(() => {
  const isSignin = false;

  return (
    <SafeAreaProvider>
      <NavigationContainer>
        <Navigator initialRouteName="introduce">
          {isSignin ? (
            <Screen name="home" component={Home} options={{ headerShown: false }} />
          ) : (
            <Screen name="signin" component={Signin} options={{ headerShown: false }} />
          )}
        </Navigator>
      </NavigationContainer>
    </SafeAreaProvider>
  );
});
