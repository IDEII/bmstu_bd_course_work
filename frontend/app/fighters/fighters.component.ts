import { Component } from '@angular/core';
import { Fighter } from '../fighter';
import { FighterService } from '../fighter.service';

@Component({
  selector: 'app-fighters',
  templateUrl: './fighters.component.html',
  styleUrl: './fighters.component.css'
})
export class FightersComponent {
  fighters: Array<Fighter>;
   
  constructor(private fighterService: FighterService) {
    this.fighters = new Array<Fighter>();
}
  ngOnInit() {
    this.fighterService.getFighters().subscribe(data => {
      this.fighters = data.filter(fighter => fighter.id !== '00000000-0000-0000-0000-000000000000');
    });
  }
}
