package wepdf;

import java.io.File;
import java.io.InputStream;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;
import java.time.Duration;
import org.json.JSONObject;

public class wepdf_render {
	private static final String API_KEY = "YOUR_API_KEY";

    public static void main(String... args) throws Exception {
        HttpRequest httpRequest = HttpRequest.newBuilder()
                .uri(URI.create("https://api.wepdf.io/v1/render?url=https://google.com&apikey="+API_KEY))
                .header("Content-Type", "application/json")
                .build();

        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_1_1)
                .build();

        HttpResponse<InputStream> response = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofInputStream());

        int statusCode = response.statusCode();
        if (statusCode == 200 || statusCode == 201) {
            // Save the file locally
            File targetFile = new File("google.pdf");
            Files.copy(response.body(), targetFile.toPath(), StandardCopyOption.REPLACE_EXISTING);
            System.out.println("success");
        } else {
            // error occurred
        	System.out.println("error");
        }
    }
}
