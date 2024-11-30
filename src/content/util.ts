export async function getPermissions() {
    const permissions = ['geolocation', 'notifications', 'microphone', 'camera']
    const result = {}
    for (const permissionName of permissions) {
      try {
        const permissionStatus = await navigator.permissions.query({
          name: permissionName
        })
        result[permissionName] = permissionStatus.state
      } catch (error) {
        result[permissionName] = 'Not available'
      }
    }
    return result
  }