/**
 *@Desc:
 *@Author:Giousa
 *@Date:2021/1/12
 */
package models

var TextController =
`package {{.Package}};


/**
 * Description:
 * Author:{{.Author}}
 * Email:{{.Email}}
 * Date:{{.DateTime}}
 */
@RestController
@RequestMapping("{{.NameSeparate}}")
public class {{.Name}}Controller {

	@Autowired
    private {{.Name}}Service {{.NameHumpLower}}Service;

    @PostMapping("add{{.Name}}")
    public {{.ResultBody}} add{{.Name}}(@RequestBody {{.Name}} {{.NameHumpLower}}) {
        return {{.NameHumpLower}}Service.add{{.Name}}({{.NameHumpLower}});
    }

	@PostMapping("update{{.Name}}")
    public {{.ResultBody}} update{{.Name}}(@RequestBody {{.Name}} {{.NameHumpLower}}) {
        return {{.NameHumpLower}}Service.update{{.Name}}({{.NameHumpLower}});
    }

    @GetMapping("find{{.Name}}ById")
    public {{.ResultBody}} find{{.Name}}ById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return {{.NameHumpLower}}Service.find{{.Name}}ById(id);
    }

    @GetMapping("delete{{.Name}}ById")
    public {{.ResultBody}} delete{{.Name}}ById(@RequestParam("id") String id) {
		if(StringUtils.isEmpty(id)){
            return ResultVO.error(ResultEnum.PARAM_EMPTY);
        }
        return {{.NameHumpLower}}Service.delete{{.Name}}ById(id);
    }

 	@GetMapping("find{{.Name}}ListByPage")
    public {{.ResultBody}} find{{.Name}}ListByPage(
                                    @RequestParam(value = "page",required = false,defaultValue = "1") int page,
                                    @RequestParam(value = "size",required = false,defaultValue = "10") int size) {
        return {{.NameHumpLower}}Service.find{{.Name}}ListByPage(page,size);
    }

}
`

var TextService =
	`package {{.Package}}.service;



/**
 * Description:
 * Author:{{.Author}}
 * Email:{{.Email}}
 * Date:{{.DateTime}}
 */
public interface {{.Name}}Service {

    {{.ResultBody}} add{{.Name}}({{.Name}} {{.NameHumpLower}});

    {{.ResultBody}} update{{.Name}}({{.Name}} {{.NameHumpLower}});

    {{.ResultBody}} find{{.Name}}ById(String id);

    {{.ResultBody}} delete{{.Name}}ById(String id);

    {{.ResultBody}} find{{.Name}}ListByPage(int page,int size);
}

`

var TextServiceImpl =
	`package {{.Package}}.service.impl;



/**
 * Description:
 * Author:{{.Author}}
 * Email:{{.Email}}
 * Date:{{.DateTime}}
 */
@Service
public class {{.Name}}ServiceImpl implements {{.Name}}Service {

    @Autowired
    private {{.Name}}Mapper {{.NameHumpLower}}Mapper;

    @Override
    public {{.ResultBody}} add{{.Name}}({{.Name}} {{.NameHumpLower}}) {

		return null;
    }

	@Override
    public {{.ResultBody}} update{{.Name}}({{.Name}} {{.NameHumpLower}}) {
		
		return null;
    }

    @Override
    public {{.ResultBody}} find{{.Name}}ById(String id) {

		return null;
    }

    @Override
    public {{.ResultBody}} delete{{.Name}}ById(String id) {
        
		return null;
    }

    @Override
    public {{.ResultBody}} find{{.Name}}ListByPage(int page,int size) {

		return null;
    }
}

`