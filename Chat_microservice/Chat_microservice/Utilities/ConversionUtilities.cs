

using System.Text;
using Newtonsoft.Json;

namespace Chat_microservice.Utilities
{
    public class ConversionUtilities
    {
        public static string SerializeJson(object obj) => JsonConvert.SerializeObject(obj);

        public static byte[] SerializeBinary(object obj)
        {
            var str = SerializeJson(obj);
            return Encoding.UTF8.GetBytes(str);
        }
        public static T DeserializeJson<T>(string json) => JsonConvert.DeserializeObject<T>(json);

        public static T DeserializeBinary<T>(byte[] data)
        {
            var jsonMessage = Encoding.UTF8.GetString(data);
            return JsonConvert.DeserializeObject<T>(jsonMessage);
        }
    }
}
