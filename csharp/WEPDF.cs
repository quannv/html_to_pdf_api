using System;
using RestSharp;
using Newtonsoft.Json;
using System.IO;
using RestSharp.Authenticators;
using RestSharp.Serialization;
using System.Net.Mail;
using System.Net;
using System.Collections.Generic;
using Newtonsoft.Json.Linq;

namespace WepdfExample
{
    class Program
    {
        static void Main(string[] args)
        {
            IRestClient client = new RestClient("http://localhost:9000/render");
            client.Authenticator = new HttpBasicAuthenticator("your_api_key", "");

            IRestRequest request = new RestRequest(Method.GET);

            var json = new
            {
                url = "https://en.wikipedia.org/wiki/PDF"
            };
            request.AddJsonBody(json);

            IRestResponse response = client.Execute(request);
            if (!response.IsSuccessful)
            {
                // Check why status is not int 2xx.
            }
            else
            {
                File.WriteAllBytes("wikipedia.pdf", response.RawBytes);
            }
        }
    }
}