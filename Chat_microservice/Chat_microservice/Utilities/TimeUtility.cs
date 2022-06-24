using System;

namespace Chat_microservice.Utilities
{
    public class TimeUtility
    {
        public static long GetUnixTimeStamp(DateTime time)
        {
            return (time.ToLocalTime().Ticks - new DateTime(1970, 1, 1).Ticks) / TimeSpan.TicksPerSecond;
        }
    }
}
