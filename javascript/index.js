const fs = require('fs');

const wepdf = require('wepdf')('YOUR_API_KEY');

const content = fs.readFileSync('your_file.html', {encoding: 'utf-8'});



//GET request
let getOptions = {
    landscape: true,
};

return wepdf.prepare(getOptions)
            .render('https://www.google.com/')
            .then(res=>{
                try{
                    fs.writeFileSync('wepdf_render.pdf',res.body);
                }catch (err) {
                    console.log("ERROR upload file===>",err);
                    throw (err)
                }
            });


//POST request
let myHeader = {
    'Content-Type': 'text/html'
};
let postOptions = {
    landscape: false,
};
return wepdf
    .prepare(postOptions)
    .setHeaders(myHeader)
    .convert(content)
    .then(res => {
        fs.writeFileSync('wepdf_convert.pdf', res.body);
    }).catch(e => {
        // console.log("eeee",e);
    });

