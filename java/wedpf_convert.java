package wepdf;

import java.io.File;
import java.io.FileInputStream;
import java.io.InputStream;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;
import java.nio.file.Files;
import java.nio.file.Paths;
import java.nio.file.StandardCopyOption;

public class wedpf_convert {
	private static final String API_KEY = "YOUR_API_KEY";

    public static void main(String... args) throws Exception {
//    	Read file as String
    	FileInputStream fis = new FileInputStream("your_file.html");
    	byte[] buffer = new byte[10];
    	StringBuilder sb = new StringBuilder();
    	while (fis.read(buffer) != -1) {
    		sb.append(new String(buffer));
    		buffer = new byte[10];
    	}
    	fis.close();

    	String content = sb.toString();
    	
    	
//    	Send content to service API
        HttpRequest httpRequest = HttpRequest.newBuilder()
                .uri(URI.create("https://api.wepdf.io/v1/render?landscape=true&apikey="+API_KEY))
                .header("Content-Type", "text/html")
                .POST(HttpRequest.BodyPublishers.ofString(content))
                .build();

        HttpClient httpClient = HttpClient.newBuilder()
                .version(HttpClient.Version.HTTP_1_1)
                .build();

        HttpResponse<InputStream> response = httpClient.send(httpRequest, HttpResponse.BodyHandlers.ofInputStream());

        int statusCode = response.statusCode();
        if (statusCode == 200 || statusCode == 201) {
            // Save the file locally
            File targetFile = new File("wepdf.pdf");
            Files.copy(response.body(), targetFile.toPath(), StandardCopyOption.REPLACE_EXISTING);
            System.out.println("success");
        } else {
            // error occurred
        	System.out.println("error");
        }
    }
}
